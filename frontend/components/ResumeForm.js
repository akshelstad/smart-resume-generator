import { useState } from "react";
import axios from "axios";

export default function ResumeForm({ setResume }) {
    const [formData, setFormData] = useState({ name: "", experience: "" });
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState("");

    const API_URL = process.env.NEXT_PUBLIC_BACKEND_URL || "http://localhost:8080";

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);
        setError("");

        try {
            const response = await axios.post(`${API_URL}/generate-resume`, formData);
            setResume(response.data.resume);
        } catch (error) {
            console.error("Error generating resume:", error.response?.data || error.message);
            setError("Failed to generate resume. Please try again.");
        }

        setLoading(false);
    };

    return (
        <form onSubmit={handleSubmit} className="bg-white p-6 shadow-md rounded-lg space-y-4">
            <input
                type="text"
                name="name"
                placeholder="Full Name"
                value={formData.name}
                onChange={handleChange}
                className="border p-2 w-full rounded"
                required
            />
            <textarea
                name="experience"
                placeholder="Describe your experience"
                value={formData.experience}
                onChange={handleChange}
                className="border p-2 w-full rounded"
                required
            />
            {error && <p className="text-red-500">{error}</p>}
            <button 
                type="submit" 
                className={`bg-blue-500 text-white p-2 rounded w-full ${!formData.name || !formData.experience ? "opacity-50 cursor-not-allowed" : ""}`}
                disabled={!formData.name || !formData.experience || loading}
            >
                {loading ? "Generating..." : "Generate Resume"}
            </button>
        </form>
    );
}
