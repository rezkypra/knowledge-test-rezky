import StatsCard from "@/components/StatsCard";
import StudentTable from "@/components/StudentTable";
import { mockStats, mockStudents } from "@/lib/mockData";
import { Users, UserPlus, FileText } from "lucide-react";

export default function DashboardPage() {
    return (
        <div className="space-y-8">
            {/* Header Section */}
            <div className="bg-white p-8 rounded-2xl shadow-sm">
                <h1 className="text-2xl font-bold text-gray-900">
                    Selamat Datang, Admin
                </h1>
                <p className="text-gray-500 mt-1">
                    Track, and manage your student&apos;s information here.
                </p>

                {/* Stats Grid */}
                <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mt-8">
                    <StatsCard
                        title="Peningkatan Mahasiswa Baru"
                        value={`${mockStats.newStudents}%`}
                        trend="vs last month"
                        trendUp={true}
                        color="bg-white"
                    />
                    <StatsCard
                        title="Total Mahasiswa"
                        value={mockStats.totalStudents}
                        icon={Users}
                        color="bg-white"
                    />
                    <StatsCard
                        title="Mahasiswa Terdaftar"
                        value={mockStats.activeStudents}
                        icon={UserPlus}
                        color="bg-white"
                    />
                </div>
            </div>

            {/* Table Section */}
            <div className="space-y-4">
                <div className="flex justify-between items-center">
                    <input
                        type="text"
                        placeholder="Search"
                        className="px-4 py-2 border border-gray-200 rounded-lg w-64 focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />
                </div>
                <StudentTable students={mockStudents} />
            </div>
        </div>
    );
}
