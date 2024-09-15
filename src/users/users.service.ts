import { Injectable, NotFoundException } from '@nestjs/common';
import { UsersModel } from './users.model';

@Injectable()
export class UsersService {
  usersList: UsersModel[] = [];

  insertUser(name: string, degree: string, age: number) {
    const user = new UsersModel(Math.random(), name, degree, age);
    this.usersList.push(user);
  }

  getUsers() {
    return [...this.usersList];
  }
  getSingleUser(userId: string) {
    return this.usersList[this.findUser(userId)];
  }
  updateUser(userId: string, name: string, degree: string, age: number) {
    const updateUser = this.usersList[this.findUser(userId)];
    if (name) {
      updateUser.name = name;
    }
    if (degree) {
      updateUser.degree = degree;
    }
    if (age) {
      updateUser.age = age;
    }
    this.usersList[this.findUser(userId)] = updateUser;
  }
  deleteUser(userId: string) {
    this.usersList.splice(this.findUser(userId), 1);
  }

  findUser(userId: string) {
    const userIndex = this.usersList.findIndex(
      (user) => user.id.toString() === userId,
    );
    if (!userIndex && userIndex !== 0) {
      throw new NotFoundException('user not found');
    }
    return userIndex;
  }
}
