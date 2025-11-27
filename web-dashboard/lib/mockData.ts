export interface Student {
    id: string;
    name: string;
    nik: string;
    joinDate: string;
    gender: "Laki-laki" | "Perempuan";
    address: string;
}

export interface Stats {
    totalStudents: number;
    activeStudents: number;
    newStudents: number;
}

export const mockStats: Stats = {
    totalStudents: 400,
    activeStudents: 20,
    newStudents: 40, // 40% increase
};

export const mockStudents: Student[] = [
    {
        id: "1",
        name: "Circoales",
        nik: "12345678",
        joinDate: "2020",
        gender: "Laki-laki",
        address: "Super lightweight design app",
    },
    {
        id: "2",
        name: "SMAN 1",
        nik: "12345678",
        joinDate: "2020",
        gender: "Perempuan",
        address: "AI and machine learning data",
    },
    {
        id: "3",
        name: "Command+R",
        nik: "12345678",
        joinDate: "2020",
        gender: "Perempuan",
        address: "Brings all your news into one place",
    },
    {
        id: "4",
        name: "Hourglass",
        nik: "12345678",
        joinDate: "2020",
        gender: "Laki-laki",
        address: "Time management and productivity",
    },
    {
        id: "5",
        name: "Layers",
        nik: "12345678",
        joinDate: "2020",
        gender: "Laki-laki",
        address: "Connect web apps seamlessly",
    },
    {
        id: "6",
        name: "Quotient",
        nik: "12345678",
        joinDate: "2020",
        gender: "Perempuan",
        address: "Web-based sales doc management",
    },
    {
        id: "7",
        name: "Sisyphus",
        nik: "12345678",
        joinDate: "2020",
        gender: "Laki-laki",
        address: "Time tracking, invoicing and expenses",
    },
];
