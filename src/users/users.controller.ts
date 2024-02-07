import { Controller, Post, Body } from "@nestjs/common";
import { UsersService } from "./users.service";

@Controller('users')

export class UsersController{
    constructor(private usersService: UsersService){}
    @Post()
    addUser(
    @Body('name') userName: string,
    @Body('degree') userDegree: string,
    @Body('age') userAge: number,
    ){
        this.usersService.insertUser(userName, userDegree, userAge);
        return { message: 'user registered successfully' };
    }
}