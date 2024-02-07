import { Injectable } from '@nestjs/common';
import { UsersModel } from "./users.model";

@Injectable()
export class UsersService{
    usersList: UsersModel[] = [];

    insertUser(name: string, degree: string, age: number){
        const user = new UsersModel(Math.random(), name, degree, age);
        this.usersList.push(user);
    }
}