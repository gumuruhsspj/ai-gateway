<?php

namespace App\Http\Controllers;

use App\Models\Division;
use Illuminate\Http\Request;

class DivisionController extends Controller
{
    // Mengambil semua daftar divisi
    public function index()
    {
        return response()->json(Division::all());
    }

    // Menyimpan data divisi baru (Context, Rules, Tone)
    public function store(Request $request)
    {
        $validated = $request->validate([
            'name'    => 'required|string',
            'slug'    => 'required|string|unique:divisions',
            'context' => 'required|string',
            'rules'   => 'required|string',
            'tone'    => 'required|in:formal,casual',
            'goal'    => 'nullable|string'
        ]);

        $division = Division::create($validated);
        return response()->json($division, 201);
    }
}