import { useState } from "react";
import type { User } from "../dto/User";
import Navigation from "../components/Navigation";

interface DashboardProps {
    user: User,
    onLogout: () => void,
}

export default function Dashboard({ user, onLogout }: DashboardProps) {
    const [activeTab, setActiveTab] = useState<'profile' | 'explore' | 'my_courses'>('profile');
    return (
        <div className="w-full bg-slate-800 p-6 rounded-2xl shadow-xl space-y-4">
            <Navigation user={user} onLogout={onLogout} activeTab={activeTab} onChangeTab={(tab) => setActiveTab(tab as 'profile' | 'explore' | 'my_courses')} />
            
            {activeTab === 'profile' && ( <div/> )}
        </div>
    );
}