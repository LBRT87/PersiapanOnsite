import { useState } from "react";
import { req } from "../api/api";
import type { LoginResponse, User } from "../dto/User";
import { useGoogleLogin } from "@react-oauth/google";

interface LoginProps{
    onLoginSuccess: (user: User) => void;
    onNavigate: (pages: 'login' | 'register' | 'dto') => void;
}

export default function Login({ onLoginSuccess }: LoginProps){
    const [, setError] = useState('');

    const handleSubmit = async(e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setError('');
        const formData = new FormData(e.currentTarget);
        const fields = {
            email: formData.get('email') as string,
            password: formData.get('password') as string
        }
        try{
            const data = await req<LoginResponse>('/auth/login', fields);

            localStorage.setItem('access_token', data.access_token);
            localStorage.setItem('refresh_token', data.refresh_token);
            localStorage.setItem('user_role', data.user.role);

            onLoginSuccess(data.user);
        } catch (err) {
            setError(err instanceof Error ? err.message : "error");
        }
    }

    const loginGoogle = useGoogleLogin({
        onSuccess: async () => {
            onerror('');
            try{
                const data = await req<LoginResponse>('/auth/google/login', { token:'access_token '});
                    
                localStorage.setItem('access_token', data.access_token);
                localStorage.setItem('refresh_token', data.refresh_token);
                localStorage.setItem('user_role', data.user.role);

                onLoginSuccess(data.user);
            } catch (err) {
                setError(err instanceof Error ? err.message : "error");
            }
        },
        onError: () => setError('Google login failed')
    })
    
    return (
        <div className="max-w-sm mx-auto bg-slate-800 p-6 rounded-xl border border-slate-700">
            <h2 className="text-xl font-bold text-white mb-4 text-center">Login to App</h2>

            <form className="space-y-4" onSubmit={handleSubmit}>
                <input
                type="email"
                placeholder="Email"
                className="w-full p-2 bg-slate-900 border border-slate-700 rounded text-white focus:border-emerald-500 outline-none"
                />
                <input
                type="password"
                placeholder="Password"
                className="w-full p-2 bg-slate-900 border border-slate-700 rounded text-white focus:border-emerald-500 outline-none"
                />

                <label className="flex items-center text-sm text-slate-300">
                <input type="checkbox" className="mr-2 accent-emerald-500" /> Keep me signed in
                </label>

                <button className="w-full py-2 bg-emerald-600 hover:bg-emerald-500 text-white font-medium rounded" onClick={()=>handleSubmit}>
                Log In
                </button>
            </form>

            <div className="my-4 text-center text-sm text-slate-500">or</div>

            <button onClick={() => loginGoogle} className="w-full py-2 bg-white text-slate-900 font-medium rounded flex items-center justify-center gap-2">
                <span>Google Login</span>
            </button>
        </div>
    );
}