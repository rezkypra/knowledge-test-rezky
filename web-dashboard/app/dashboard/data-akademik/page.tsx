import { Award, TrendingUp, AlertCircle } from "lucide-react";

export default function DataAkademikPage() {
    return (
        <div className="space-y-8">
            <div>
                <h1 className="text-2xl font-bold text-gray-900">Data Akademik</h1>
                <p className="text-gray-500">Overview of academic performance.</p>
            </div>

            {/* Performance Overview */}
            <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
                <div className="bg-gradient-to-br from-blue-500 to-blue-600 p-6 rounded-xl text-white shadow-lg">
                    <div className="flex items-center gap-3 mb-4">
                        <Award className="opacity-80" />
                        <span className="font-medium opacity-90">Rata-rata IPK</span>
                    </div>
                    <div className="text-4xl font-bold mb-2">3.45</div>
                    <div className="text-sm opacity-80">↑ 0.12 dari semester lalu</div>
                </div>

                <div className="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
                    <div className="flex items-center gap-3 mb-4 text-gray-900">
                        <TrendingUp className="text-green-500" />
                        <span className="font-medium">Kelulusan Tepat Waktu</span>
                    </div>
                    <div className="text-4xl font-bold mb-2 text-gray-900">85%</div>
                    <div className="text-sm text-gray-500">Target: 90%</div>
                </div>

                <div className="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
                    <div className="flex items-center gap-3 mb-4 text-gray-900">
                        <AlertCircle className="text-orange-500" />
                        <span className="font-medium">Mahasiswa Cuti</span>
                    </div>
                    <div className="text-4xl font-bold mb-2 text-gray-900">12</div>
                    <div className="text-sm text-gray-500">Perlu konfirmasi ulang</div>
                </div>
            </div>

            {/* Recent Academic Activities */}
            <div className="bg-white rounded-xl shadow-sm border border-gray-100 p-6">
                <h3 className="text-lg font-bold text-gray-900 mb-4">
                    Aktivitas Akademik Terbaru
                </h3>
                <div className="space-y-4">
                    {[1, 2, 3].map((i) => (
                        <div
                            key={i}
                            className="flex items-center gap-4 p-4 bg-gray-50 rounded-lg"
                        >
                            <div className="w-10 h-10 rounded-full bg-blue-100 flex items-center justify-center text-blue-600 font-bold">
                                IF
                            </div>
                            <div className="flex-1">
                                <h4 className="font-medium text-gray-900">
                                    Input Nilai Ujian Tengah Semester
                                </h4>
                                <p className="text-sm text-gray-500">
                                    Dosen Pengampu: Dr. Budi Santoso
                                </p>
                            </div>
                            <span className="text-sm text-gray-400">2 jam yang lalu</span>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
}
