import { useState } from "react";
import Login from "./pages/Login";
import type { User } from "./dto/User";
import { GoogleOAuthProvider } from "@react-oauth/google";

export default function App() {
  const [page, setPage] = useState<'login' | 'register' | 'otp' | 'dashboard'>('login');
  // const [email, setEmail] = useState('');
  const [, setUser] = useState<User | null>(null);
  // const [, setUser] = useState<User | null>({
  //   id: 1,
  //   email: "wilson@gmail.com",
  //   username: "wilson",
  //   password: "wilson",
  //   role: "student",
  //   dob: "2007-06-25"
  // });
  
  // const handleLogout = () => {
  //   localStorage.clear();
  //   setUser(null);
  //   setPage('login');
  // }

  return (
    // <GoogleOAuthProvider clientId="">
      <div className={`min-h-screen bg-slate-900 text-white p-4 flex ${
        page === 'dashboard' ? 'flex-col items-stretch' : 'items-center justify-center'
      }`}>
        { page === 'login' && (<Login onLoginSuccess={(u) => { setUser(u); setPage('dashboard'); }} onNavigate={() => setPage}/>) }
        {/* { page === 'register' && (<Register onNavigate={setPage} onOtpSent={(em) => { setEmail(em); setPage('otp'); }}/>) }
        { page === 'otp' && (<VerifyOtp email={email} onSuccess={ () => setPage('login') }/>) }
        { page === 'dashboard' && user && (<Dashboard user={user} onLogout={handleLogout}/>)} */}
      </div>
    // </GoogleOAuthProvider>
  );
}