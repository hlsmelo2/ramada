<?php

use App\Http\Controllers\ProductsController;
use App\Http\Controllers\UsersController;
use Illuminate\Support\Facades\Route;

Route::group(['middleware' => 'AuthGuard'], function() {
    Route::get('/', [ProductsController::class, 'pageList'])->name('home');
    Route::get('/login', [UsersController::class, 'pageLogin'])->name('login');
    Route::post('/login', [UsersController::class, 'login'])->name('try.login');
    Route::get('/logout', [UsersController::class, 'logout'])->name('logout');

    Route::group(['prefix' => '/produtos'], function() {
        Route::get('', [ProductsController::class, 'pageList'])->name('products');
        Route::get('/create', [ProductsController::class, 'pageCreate'])->name('page.product.create');
        Route::get('/{id}/edit', [ProductsController::class, 'pageEdit'])->name('page.product.edit');
        Route::get('/{id}/delete', [ProductsController::class, 'pageDelete'])->name('page.product.delete');
        Route::post('', [ProductsController::class, 'edit'])->name('product.create');
        Route::put('/{id}', [ProductsController::class, 'edit'])->name('product.edit');
        Route::delete('/{id}', [ProductsController::class, 'delete'])->name('product.delete');
    });

    Route::group(['prefix' => '/usuarios'], function() {
        Route::get('', [UsersController::class, 'pageList'])->name('users');
        Route::get('/create', [UsersController::class, 'pageCreate'])->name('page.user.create');
        Route::get('/{id}/edit', [UsersController::class, 'pageEdit'])->name('page.user.edit');
        Route::get('/{id}/delete', [UsersController::class, 'pageDelete'])->name('page.user.delete');
        Route::post('', [UsersController::class, 'edit'])->name('user.create');
        Route::put('/{id}', [UsersController::class, 'edit'])->name('user.edit');
        Route::delete('/{id}', [UsersController::class, 'delete'])->name('user.delete');
    });
});
