<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use Symfony\Component\HttpFoundation\Response;

class AuthGuard
{
    /**
     * Handle an incoming request.
     *
     * @param  \Closure(\Illuminate\Http\Request): (\Symfony\Component\HttpFoundation\Response)  $next
     */
    public function handle(Request $request, Closure $next): Response
    {
        $token = $_COOKIE['_token'] ?? null;
        $allowedPage = in_array(Route::currentRouteName(), ['login', 'try.login', 'page.user.create', 'user.create']);
        $authed = $token !== null;

        if (!$authed) {
            if (!$allowedPage) {
                return redirect()->route('login');
            }

            return $next($request);
        }

        if (Route::currentRouteName() === 'login') {
            return redirect()->route('products');
        }

        return $next($request);
    }
}
