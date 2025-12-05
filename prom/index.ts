import express from "express";
import client from "prom-client"
import type { NextFunction, Request, Response } from "express";

const app = express();

const requestHanlder = new client.Counter({
    name: "http_total_requests",
    help: "Total number of request",
    labelNames: ['method', 'route', 'status_code']
})

const gaugeHandler = new client.Gauge({
    name: "active_requests",
    help: "Number of active request"
})

function middleware(req: Request, res: Response, next: NextFunction){
    const startTime = Date.now();
    gaugeHandler.inc();
    res.on('finish', () => {
        const endTime = Date.now();
        console.log(`Request took ${endTime - startTime}ms`);

        requestHanlder.inc({
            method: req.method,
            route: req.route? req.route.path: req.path,
            status_code: res.statusCode
        });
        gaugeHandler.dec();
    })
    next()
}

app.use(middleware)

app.get('/cpu', (req, res) => {
    for(let i = 0; i < 10000000; i++){
        Math.random();
    }

    res.json({
        message: "Cpu"
    })
})


app.get('/health', (req, res) => {
    res.json({
        msg: "hello"
    })
})

app.get('/metrics', async (req, res) => {
    const metrics = await client.register.metrics();
    res.set('Content-Type', client.register.contentType);
    res.end(metrics);
})

app.listen(3000, () => {
    console.log("App started successfully");
})