import { Student } from "@/lib/mockData";
import { Pencil, Trash2 } from "lucide-react";

interface StudentTableProps {
    students: Student[];
}

export default function StudentTable({ students }: StudentTableProps) {
    return (
        <div className="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden">
            <div className="overflow-x-auto">
                <table className="w-full text-sm text-left">
                    <thead className="text-xs text-gray-500 uppercase bg-gray-50 border-b">
                        <tr>
                            <th className="px-6 py-3 font-medium">Nama Mahasiswa</th>
                            <th className="px-6 py-3 font-medium">NIK</th>
                            <th className="px-6 py-3 font-medium">Masuk pada</th>
                            <th className="px-6 py-3 font-medium">Jenis Kelamin</th>
                            <th className="px-6 py-3 font-medium">Alamat</th>
                            <th className="px-6 py-3 font-medium text-right">Action</th>
                        </tr>
                    </thead>
                    <tbody>
                        {students.map((student) => (
                            <tr
                                key={student.id}
                                className="bg-white border-b hover:bg-gray-50 transition-colors"
                            >
                                <td className="px-6 py-4 font-medium text-gray-900 flex items-center gap-3">
                                    <div className="w-8 h-8 rounded-full bg-gray-200 flex-shrink-0"></div>
                                    {student.name}
                                </td>
                                <td className="px-6 py-4 text-gray-500">{student.nik}</td>
                                <td className="px-6 py-4">
                                    <span className="bg-green-100 text-green-800 text-xs font-medium px-2.5 py-0.5 rounded">
                                        {student.joinDate}
                                    </span>
                                </td>
                                <td className="px-6 py-4 text-gray-500">{student.gender}</td>
                                <td className="px-6 py-4 text-gray-500 truncate max-w-xs">
                                    {student.address}
                                </td>
                                <td className="px-6 py-4 text-right">
                                    <div className="flex items-center justify-end gap-2">
                                        <button className="p-1 text-gray-400 hover:text-red-500">
                                            <Trash2 size={16} />
                                        </button>
                                        <button className="p-1 text-gray-400 hover:text-blue-500">
                                            <Pencil size={16} />
                                        </button>
                                    </div>
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
            <div className="p-4 border-t border-gray-100 flex justify-between items-center">
                <span className="text-sm text-gray-500">Page 1 of 10</span>
                <div className="flex gap-1">
                    <button className="px-3 py-1 border rounded hover:bg-gray-50 text-sm">
                        &lt;
                    </button>
                    <button className="px-3 py-1 border rounded bg-blue-50 text-blue-600 text-sm">
                        1
                    </button>
                    <button className="px-3 py-1 border rounded hover:bg-gray-50 text-sm">
                        2
                    </button>
                    <button className="px-3 py-1 border rounded hover:bg-gray-50 text-sm">
                        &gt;
                    </button>
                </div>
            </div>
        </div>
    );
}
