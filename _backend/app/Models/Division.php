<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Division extends Model
{
    use HasFactory;

    protected $fillable = [
        'name',
        'slug',
        'context',
        'rules',
        'tone',
        'goal',
    ];

    /**
     * Opsional: Jika ingin slug selalu lowercase otomatis
     */
    public function setSlugAttribute($value)
    {
        $this->attributes['slug'] = strtolower($value);
    }
    
}
