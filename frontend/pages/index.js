import { useState } from "react";
import ResumeForm from "../components/ResumeForm";
import ResumePreview from "../components/ResumePreview";

export default function Home() {
    const [resume, setResume] = useState("");

    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-50">
            <div className="w-full max-w-lg">
                <h1 className="text-2xl font-bold text-center mb-4">AI Resume Generator</h1>
                <ResumeForm setResume={setResume} />
                {resume ? <ResumePreview resume={resume} /> : <p className="text-center text-gray-500">Enter details to generate a resume.</p>}
            </div>
        </div>
    );
}


