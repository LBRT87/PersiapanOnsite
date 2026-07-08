import type { User } from "../dto/User";

interface NavigationProps{
    user: User;
    onLogout: () => void;
    activeTab?: string;
    onChangeTab?: (tab: string) => void;
}

export default function Navigation({user, onLogout}: NavigationProps){
    return(
        <nav className="bg-slate-900 border-b border-slate-800 p-4">
            <div className="max-w-7xl mx-auto flex items-center justify-between">

                <div className="flex items-center gap-6">
                <h1 className="text-emerald-500 font-bold text-xl">App</h1>
                <div className="hidden md:flex gap-4 text-sm text-slate-300">
                    <a href="#" className="hover:text-emerald-400">Explore</a>
                    <a href="#" className="hover:text-emerald-400">My Courses</a>
                    <a href="#" className="hover:text-emerald-400">Messages</a>
                </div>
                </div>

                <div className="flex items-center gap-4 text-sm">
                <button className="text-slate-300">🛒 Cart</button>
                <div className="border-l border-slate-700 pl-4 flex items-center gap-2">
                    <div className="w-8 h-8 bg-slate-700 rounded-full text-center leading-8 text-white">{user.username}</div>
                    <button className="text-red-400 hover:underline" onClick={onLogout}>Logout</button>
                </div>
                </div>

            </div>
        </nav>
    );
}