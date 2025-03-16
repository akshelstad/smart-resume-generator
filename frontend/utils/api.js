import axios from "axios";

const API_URL = process.env.NEXT_PUBLIC_BACKEND_URL || "http://localhost:8080";

export const generateResume = async (data) => {
    try {
        const response = await axios.post(`${API_URL}/generate-resume`, data);
        return response.data?.resume || "No resume generated.";
    } catch (error) {
        console.error("Error generating resume:", error.response?.data || error.message);
        return { error: error.response?.data?.message || "An unexpected error occurred" };
    }
};
