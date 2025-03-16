export default function ResumePreview({ resume }) {
    return (
        <div className="bg-gray-100 p-6 shadow-md rounded-lg mt-4">
            <h2 className="text-lg font-semibold">Generated Resume</h2>
            <p className="whitespace-pre-line">{resume}</p>
        </div>
    );
}