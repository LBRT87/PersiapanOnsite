import { useState } from "react";
import type { RegisterResponse } from "../dto/User";
import { req } from "../api/Api";


interface RegisterProps{
    onOtpSent: (email: string) => void
    onNavigate: (page: 'register' | 'login' | 'otp') => void;
}

export default function Register({ onOtpSent, onNavigate }: RegisterProps){
    const [, setError] = useState('');

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setError('');
        const formData = new FormData(e.currentTarget);
        const fields = {
            username: formData.get('username') as string,
            date: formData.get('dob') as string,
            email: formData.get('email') as string,
            password: formData.get('password') as string,
            role: 'student'
        }
        try {
            await req<RegisterResponse>('/auth/register', fields);
            onOtpSent(fields.email as string);
            onNavigate('otp');
        } catch (err) {
            setError(err instanceof Error ? err.message : "error"); 
        }
    }

    return (
        <div className="max-w-md mx-auto bg-slate-800 p-6 rounded-xl border border-slate-700">
            <h2 className="text-xl font-bold text-white mb-4 text-center">Create Account</h2>

            <form className="space-y-4 flex flex-col" onSubmit={handleSubmit}>
                <div className="grid grid-cols-2 gap-4">
                <input type="text" placeholder="Username" className="p-2 bg-slate-900 border border-slate-700 rounded text-white" />
                <input type="date" className="p-2 bg-slate-900 border border-slate-700 rounded text-white" />
                </div>

                <input type="email" placeholder="Email Address" className="p-2 bg-slate-900 border border-slate-700 rounded text-white" />
                <input type="password" placeholder="Password" className="p-2 bg-slate-900 border border-slate-700 rounded text-white" />
                <input type="password" placeholder="Confirm Password" className="p-2 bg-slate-900 border border-slate-700 rounded text-white" />

                <div className="p-2 border border-slate-700 rounded bg-slate-900">
                <label className="text-sm text-slate-400 block mb-1">Profile Picture</label>
                <input type="file" accept="image/*" className="text-slate-300 text-sm" />
                </div>

                <button className="py-2 bg-emerald-600 text-white font-medium rounded">
                    Register
                </button>

                <button onClick={() => onNavigate('login')} className="py-2 bg-emerald-600 text-white font-medium rounded">
                    Have an account? login here
                </button>
            </form>
        </div>
    );
}