import { BookOpen, Clock, Users } from "lucide-react";

const courses = [
    {
        id: 1,
        name: "Algoritma & Pemrograman",
        code: "IF101",
        sks: 4,
        students: 45,
        schedule: "Senin, 08:00 - 11:00",
        color: "bg-blue-50 text-blue-600",
    },
    {
        id: 2,
        name: "Struktur Data",
        code: "IF102",
        sks: 3,
        students: 38,
        schedule: "Selasa, 13:00 - 15:30",
        color: "bg-purple-50 text-purple-600",
    },
    {
        id: 3,
        name: "Basis Data",
        code: "IF201",
        sks: 3,
        students: 42,
        schedule: "Rabu, 09:00 - 11:30",
        color: "bg-green-50 text-green-600",
    },
    {
        id: 4,
        name: "Pemrograman Web",
        code: "IF202",
        sks: 4,
        students: 40,
        schedule: "Kamis, 08:00 - 11:00",
        color: "bg-orange-50 text-orange-600",
    },
];

export default function MataKuliahPage() {
    return (
        <div className="space-y-6">
            <div className="flex justify-between items-center">
                <div>
                    <h1 className="text-2xl font-bold text-gray-900">
                        Daftar Mata Kuliah
                    </h1>
                    <p className="text-gray-500">
                        Manage curriculum and course schedules.
                    </p>
                </div>
                <button className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
                    + Tambah Mata Kuliah
                </button>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                {courses.map((course) => (
                    <div
                        key={course.id}
                        className="bg-white p-6 rounded-xl shadow-sm border border-gray-100 hover:shadow-md transition-shadow"
                    >
                        <div className="flex justify-between items-start mb-4">
                            <div className={`p-3 rounded-lg ${course.color}`}>
                                <BookOpen size={24} />
                            </div>
                            <span className="text-xs font-medium text-gray-400 border px-2 py-1 rounded">
                                {course.code}
                            </span>
                        </div>
                        <h3 className="text-lg font-bold text-gray-900 mb-2">
                            {course.name}
                        </h3>
                        <div className="space-y-3 text-sm text-gray-500">
                            <div className="flex items-center gap-2">
                                <Clock size={16} />
                                <span>{course.schedule}</span>
                            </div>
                            <div className="flex items-center gap-2">
                                <Users size={16} />
                                <span>{course.students} Mahasiswa</span>
                            </div>
                            <div className="flex items-center gap-2">
                                <span className="font-medium text-gray-900">{course.sks} SKS</span>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}
