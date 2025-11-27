import { Student } from "./mockData";

// Base URL for the backend API
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

export const api = {
    // Students
    getStudents: async (): Promise<Student[]> => {
        try {
            const response = await fetch(`${API_BASE_URL}/students`);
            if (!response.ok) throw new Error("Failed to fetch students");
            return await response.json();
        } catch (error) {
            console.error("API Error:", error);
            // Fallback to mock data if API fails (for development)
            return import("./mockData").then((m) => m.mockStudents);
        }
    },

    createStudent: async (student: Omit<Student, "id">) => {
        const response = await fetch(`${API_BASE_URL}/students`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(student),
        });
        return response.json();
    },

    // Courses
    getCourses: async () => {
        const response = await fetch(`${API_BASE_URL}/courses`);
        return response.json();
    },
};
