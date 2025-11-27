import { LucideIcon } from "lucide-react";

interface StatsCardProps {
    title: string;
    value: string | number;
    icon?: LucideIcon;
    trend?: string;
    trendUp?: boolean;
    color?: string;
}

export default function StatsCard({
    title,
    value,
    icon: Icon,
    trend,
    trendUp,
    color = "bg-white",
}: StatsCardProps) {
    return (
        <div className={`p-6 rounded-xl shadow-sm border border-gray-100 ${color}`}>
            <div className="flex justify-between items-start mb-4">
                <h3 className="text-sm font-medium text-gray-500">{title}</h3>
                {Icon && <Icon className="text-blue-500" size={24} />}
            </div>
            <div className="flex items-end gap-2">
                <span className="text-3xl font-bold text-gray-900">{value}</span>
            </div>
            {trend && (
                <div
                    className={`flex items-center mt-2 text-xs font-medium ${trendUp ? "text-green-500" : "text-red-500"
                        }`}
                >
                    <span>{trend}</span>
                </div>
            )}
        </div>
    );
}
