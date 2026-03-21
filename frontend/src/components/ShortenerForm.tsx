"use client";

import React, { useState } from 'react';
import { Link, Copy, Check, ArrowRight, Loader2 } from 'lucide-react';

export default function ShortenerForm() {
    const [longUrl, setLongUrl] = useState('');
    const [shortUrl, setShortUrl] = useState('');
    const [loading, setLoading] = useState(false);
    const [copied, setCopied] = useState(false);
    const [error, setError] = useState('');

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setLoading(true);
        setError('');
        setShortUrl('');

        try {
            const apiURL = process.env.NEXT_PUBLIC_API_URL;
            const res = await fetch(`${apiURL}/shorten`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ url: longUrl }),
            })

            const data = await res.json();

            if (!res.ok) throw new Error(data.error || "Something went wrong");

            setShortUrl(data.short_url);
        } catch (err: any) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
    };

    const copyToClipboard = () => {
        navigator.clipboard.writeText(shortUrl);
        setCopied(true);
        setTimeout(() => setCopied(false), 2000);
    };

    return (
        <div className="w-full max-w-2xl mx-auto space-y-6">
            <form onSubmit={handleSubmit} className="relative group">
                <div className="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                    <Link className="h-5 w-5 text-gray-400 group-focus-within:text-blue-500 transition-colors" />
                </div>
                <input
                    type="text"
                    value={longUrl}
                    onChange={(e) => setLongUrl(e.target.value)}
                    placeholder="Paste your long link here..."
                    className="block w-full pl-12 pr-32 py-4 bg-white border border-gray-200 rounded-2xl shadow-sm focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none text-gray-700 transition-all"
                    required
                />
                <button
                    type="submit"
                    disabled={loading}
                    className="absolute right-2 top-2 bottom-2 px-6 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-xl flex items-center gap-2 transition-all disabled:opacity-50"
                >
                    {loading ? <Loader2 className="animate-spin h-4 w-4" /> : "Shorten"}
                    {!loading && <ArrowRight className="h-4 w-4" />}
                </button>
            </form>

            {/* Error Message */}
            {error && (
                <div className="p-4 bg-red-50 border border-red-100 text-red-600 rounded-xl text-sm animate-in fade-in slide-in-from-top-2">
                    {error}
                </div>
            )}

            {/* Result Card */}
            {shortUrl && (
                <div className="p-6 bg-blue-50 border border-blue-100 rounded-2xl flex items-center justify-between animate-in zoom-in-95 duration-300">
                    <div className="overflow-hidden">
                        <p className="text-sm text-blue-600 font-semibold mb-1">Your short link is ready!</p>
                        <p className="text-lg font-mono text-blue-900 truncate">{shortUrl}</p>
                    </div>
                    <button
                        onClick={copyToClipboard}
                        className={`flex items-center gap-2 px-4 py-2 rounded-lg font-medium transition-all ${copied ? "bg-green-500 text-white" : "bg-white text-blue-600 border border-blue-200 hover:bg-blue-100"
                            }`}
                    >
                        {copied ? <Check className="h-4 w-4" /> : <Copy className="h-4 w-4" />}
                        {copied ? "Copied!" : "Copy"}
                    </button>
                </div>
            )}
        </div>
    );
}