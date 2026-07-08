import { useState } from "react";
import { req } from "../api/Api";

export interface VerifyOtpProps{
    email: string;
    onSuccess: () => void;
}

export default function VerifyOtp({ email, onSuccess }: VerifyOtpProps){
    const [, setError] = useState('');

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setError('');
        try{
            const formData = new FormData(e.currentTarget);
            const code = formData.get('code') as string;
            await req<void>('auth/register/verify', {email, code})
            alert('Registration success')
            onSuccess();
        } catch (err){
            setError(err instanceof Error ? err.message : "error");
        }
    }

    return (
        <div className="max-w-sm mx-auto bg-slate-800 p-6 rounded-xl border border-slate-700">
            <h2 className="text-xl font-bold text-white mb-2 text-center">Verify OTP</h2>
            <p className="text-sm text-slate-400 mb-6 text-center">
                We have sent a 6-digit verification code to <br />
                <span className="text-emerald-400 font-semibold">user@gmail.com</span>
            </p>

            <form className="space-y-6" onSubmit={handleSubmit}>
                <div className="flex justify-between gap-2">
                {[...Array(6)].map((_, i) => (
                    <input
                    key={i}
                    type="text"
                    className="w-11 h-12 text-center bg-slate-900 border border-slate-700 rounded-lg text-white text-lg font-bold focus:border-emerald-500 outline-none transition-colors"
                    />
                ))}
                </div>

                <button className="w-full py-2.5 bg-emerald-600 hover:bg-emerald-500 text-white font-semibold rounded-lg transition-colors">
                Verify Code
                </button>
            </form>

            <div className="mt-6 text-center text-sm text-slate-400 border-t border-slate-700/50 pt-4">
                Didn't receive the code?{' '}
                <button className="text-emerald-500 hover:text-emerald-400 font-medium hover:underline">
                Resend (59s)
                </button>
            </div>
        </div>
    );
}