import type { NextFunction, Request, Response } from "express";


export async function timeTook(req: Request, res: Response, next: NextFunction){
    const timeStart = Date.now();
    next();
    const endTime = Date.now();
    console.log("Time took is: " + (endTime - timeStart));
}