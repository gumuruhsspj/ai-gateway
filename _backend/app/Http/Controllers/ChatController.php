<?php

namespace App\Http\Controllers;

use App\Models\Division;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class ChatController extends Controller
{
   public function generate(Request $request)
{
    // 1. Ambil data divisi (Context, Rules, Tone)
    $division = Division::where('slug', $request->div)->firstOrFail();

    // 2. Gabungkan Context & Question
    $prompt = "You are an AI for {$division->name}.\n" .
              "Context: {$division->context}\n" .
              "Rules: {$division->rules}\n" .
              "Question: {$request->message}";

    try {
        // 3. Tembak Groq pakai Model yang kamu tes di curl tadi
        $response = Http::withHeaders([
            'Authorization' => 'Bearer ' . env('GROQ_API_KEY'),
            'Content-Type'  => 'application/json',
        ])->timeout(30)->post('https://api.groq.com/openai/v1/chat/completions', [
            'model' => 'llama-3.1-8b-instant', // SAMA DENGAN CURL KAMU
            'messages' => [
                ['role' => 'user', 'content' => $prompt]
            ],
            'temperature' => 0.6,
        ]);

        if ($response->successful()) {
            return response()->json([
                'reply'    => $response->json('choices.0.message.content'),
                'division' => $division->slug
            ]);
        }

        // Kalau gagal, kita intip error dari Groq-nya
        return response()->json(['error' => 'Groq Error', 'msg' => $response->body()], 500);

    } catch (\Exception $e) {
        return response()->json(['error' => $e->getMessage()], 500);
    }
}
}