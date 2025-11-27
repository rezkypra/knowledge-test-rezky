"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { Users, BookOpen, GraduationCap, LayoutDashboard } from "lucide-react";
import { clsx } from "clsx";

const menuItems = [
    {
        label: "Daftar Mahasiswa",
        href: "/dashboard",
        icon: Users,
    },
    {
        label: "Daftar Mata Kuliah",
        href: "/dashboard/mata-kuliah",
        icon: BookOpen,
    },
    {
        label: "Data Akademik",
        href: "/dashboard/data-akademik",
        icon: GraduationCap,
    },
];

export default function Sidebar() {
    const pathname = usePathname();

    return (
        <aside className="w-64 bg-[#1C1E26] text-white h-screen fixed left-0 top-0 flex flex-col">
            <div className="p-6">
                <h1 className="text-xl font-bold mb-1">Sistem Mahasiswa</h1>
                <div className="border-b border-gray-700 my-4"></div>
            </div>

            <nav className="flex-1 px-4 space-y-2">
                {menuItems.map((item) => {
                    const isActive = pathname === item.href;
                    return (
                        <Link
                            key={item.href}
                            href={item.href}
                            className={clsx(
                                "flex items-center gap-3 px-4 py-3 rounded-lg transition-colors",
                                isActive
                                    ? "bg-gray-700 text-white"
                                    : "text-gray-400 hover:bg-gray-800 hover:text-white"
                            )}
                        >
                            <item.icon size={20} />
                            <span className="text-sm font-medium">{item.label}</span>
                        </Link>
                    );
                })}
            </nav>

            <div className="p-6 border-t border-gray-700">
                <div className="flex items-center gap-3">
                    <div className="w-10 h-10 rounded-full bg-gray-600 flex items-center justify-center">
                        <span className="text-xs">Admin</span>
                    </div>
                    <div>
                        <p className="text-sm font-medium">Administrator</p>
                        <p className="text-xs text-gray-400">admin@digital.com</p>
                    </div>
                </div>
            </div>
        </aside>
    );
}
