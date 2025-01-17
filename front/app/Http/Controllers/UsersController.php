<?php

namespace App\Http\Controllers;

use Exception;
use Illuminate\Hashing\BcryptHasher;
use Illuminate\Http\Request;

class UsersController extends Controller
{
    public function pageList()
    {
        $http = new RequestsController();

        $response = $http->get(getApiUrl('/users'));
        $data = $response->getBody();
        $data = json_decode($data->getContents());

        return view('pages/users/users', [
            'data' => $data,
        ]);
    }

    public function login(Request $resquest) {
        $http = new RequestsController();

        $response = $http->post(getApiUrl("/login"), [
            'json' => [
                'Email' => $resquest->get('email'),
                'Password' => $resquest->get('password'),
            ],
        ]);

        $data = $response->getBody()->getContents();
        $data = json_decode($data);

        if (!isset($data->Token)) {
            session()->flash('loginFailed', 'Error ao tentar fazer login');

            return redirect()->route('login');
        }

        setcookie("_token", $data->Token, time() + (60 * 60 * 6), '/');

        return redirect()->route('products')->with('data', []);
    }

    public function logout() {
        $http = new RequestsController();

        $response = $http->post(getApiUrl("/logout"));
        $data = $response->getBody();
        $data = json_decode($data->getContents());
        setcookie("_token", '', time() - 1000, '/');

        return redirect()->route('login');
    }

    public function pageLogin()
    {
        return view('pages/login');
    }

    public function pageCreate()
    {
        return view('pages/users/edit', [
            'mode' => 'create',
            'data' => (object)[
                'ID' => 0,
                'Name' => '',
                'Username' => '',
                'Email' => '',
                'Password' => '',
            ],
        ]);
    }

    public function pageEdit(string $id)
    {
        $http = new RequestsController();

        $response = $http->get(getApiUrl("/users/$id"));
        $data = $response->getBody();
        $data = json_decode($data->getContents());

        return view("pages/users/edit", [
            'data' => $data,
            'mode' => 'edit',
        ]);
    }

    public function pageDelete(string $id)
    {
        $http = new RequestsController();

        $response = $http->get(getApiUrl("/users/$id"));
        $data = $response->getBody();
        $data = json_decode($data->getContents());

        return view('pages/users/delete', [
            'data' => $data,
        ]);
    }

    public function edit(Request $request, int  $id = 0)
    {
        $http = new RequestsController();
        $isEdit = $request->method() === "PUT";

        $data = [
            'Name' => $request->get('name'),
            'Username' => $request->get('username'),
            'Email' => $request->get('email'),
            'Password' => $request->get('password') ?? '',
        ];

        if ($isEdit) {
            $response = $http->put(getApiUrl("/users/$id"), [ 'json' => $data ]);
        } else {
            $response = $http->post(getApiUrl("/users"), [ 'json' => $data ]);
        }

        $data = $response->getBody();
        $data = json_decode($data->getContents());

        $route = $isEdit ? 'users' : 'login';
        $flashKey = $isEdit ? 'userUpdate' : 'userCreate';
        $action = $isEdit ? 'atualizado' : 'adicionado';
        $route = isAuthed() ? 'users' : $route;

        session()->flash($flashKey, "Usuário $action com sucesso");

        return redirect()->route($route);
    }

    public function delete(string $id)
    {
        $http = new RequestsController();

        $response = $http->delete(getApiUrl("/users/$id"));
        $data = $response->getBody();
        $data = json_decode($data->getContents());

        session()->flash('userDelete', 'Usuário excluído com sucesso');

        return redirect()->route('users')->with('data', $data);
    }
}
