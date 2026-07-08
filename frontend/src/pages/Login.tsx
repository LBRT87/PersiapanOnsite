import type { User } from "../dto/User";

interface LoginProps{
    onLoginSuccess: (user: User) => void;
    onNavigate: (pages: 'login' | 'register' | 'dto') => void;
}

export default function Login({ onLoginSuccess, onNavigate }: LoginProps){
    const [error, setError] = useState('');

    const handleSubmit = async(e: React.FormEvent<HTMLInputElement>) => {
        e.preventDefault();
        setError('');
        try{

        } catch (err) {
            
        }
    }
}