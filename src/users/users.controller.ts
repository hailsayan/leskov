import {
  Controller,
  Post,
  Get,
  Body,
  Param,
  Patch,
  Delete,
} from '@nestjs/common';
import { UsersService } from './users.service';

@Controller('users')
export class UsersController {
  constructor(private usersService: UsersService) {}
  @Post()
  addUser(
    @Body('name') userName: string,
    @Body('degree') userDegree: string,
    @Body('age') userAge: number,
  ) {
    this.usersService.insertUser(userName, userDegree, userAge);
    return { message: 'user registered successfully' };
  }
  @Get()
  getAllUsers() {
    return this.usersService.getUsers();
  }
  @Get(':id')
  getSingleUser(@Param('id') userId: string) {
    return this.usersService.getSingleUser(userId);
  }
  @Patch(':id')
  updateUserInfo(
    @Param('id') userId: string,
    @Body('name') userName: string,
    @Body('degree') userDegree: string,
    @Body('age') userAge: number,
  ) {
    this.usersService.updateUser(userId, userName, userDegree, userAge);
    return { message: 'user updated successfully' };
  }
  @Delete(':id')
  removeUser(@Param('id') userId: string) {
    this.usersService.deleteUser(userId);
  }
}
