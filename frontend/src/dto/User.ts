export interface User{
    id: number;
    email: string;
    username: string;
    password: string;
    dob: string;
}

export interface LoginResponse{
    access_token: string;
    refresh_token: string;
    user: User;
}

export interface RegisterResponse{
    message: string;
}