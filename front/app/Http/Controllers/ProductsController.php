<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class ProductsController extends Controller
{
    public function getFilters(Request $request)
    {
        $filters = [];

        $data = [
            'nome' => $request->get('nome') ?? '',
            'categoria' => $request->get('categoria') ?? '',
            'preco_min' => $request->get('preco_min') ?? '',
            'preco_max' => $request->get('preco_max') ?? '',
        ];

        foreach ($data as $key => $value) {
            if ($value === '') {
                continue;
            }

            $filters[] = "$key=$value";
        }

        return '/?' . implode('&', $filters);
    }

    public function pageList(Request $request)
    {
        $http = new RequestsController();
        $filters = $this->getFilters($request);
        $response = $http->get(getApiUrl('/products' . $filters));
        $data = $response->getBody();
        $data = json_decode($data->getContents());

        return view('pages/products/products', [
            'data' => $data,
            'filters' => [
                'nome' => $request->get('nome'),
                'categoria' => $request->get('categoria'),
                'preco_min' => $request->get('preco_min'),
                'preco_max' => $request->get('preco_max'),
            ],
        ]);
    }

    public function pageCreate()
    {
        return view('pages/products/edit', [
            'mode' => 'create',
            'data' => (object)[
                'ID' => 0,
                'Name' => '',
                'Description' => '',
                'Price' => '',
                'Category' => '',
            ],
        ]);
    }

    public function pageEdit(string $id, Request $request)
    {
        $http = new RequestsController();

        $response = $http->get(getApiUrl("/products/$id"));
        $data = $response->getBody();
        $data = json_decode($data->getContents());

        return view('pages/products/edit', [
            'data' => $data,
            'mode' => 'edit',
        ]);
    }

    public function pageDelete(string $id)
    {
        $http = new RequestsController();

        $response = $http->get(getApiUrl("/products/$id"));
        $data = $response->getBody();
        $data = json_decode($data->getContents());

        return view('pages/products/delete', [
            'data' => $data,
        ]);
    }

    public function edit(int $id = 0, Request $request)
    {
        $http = new RequestsController();
        $isEdit = $request->method() === "PUT";

        $data = [
            'Name' => $request->get('name'),
            'Description' => $request->get('description'),
            'Price' => currencyFloat($request->get('price')),
            'Category' => $request->get('category'),
        ];

        if ($isEdit) {
            $response = $http->put(getApiUrl("/products/$id"), [ 'json' => $data ]);
        } else {
            $response = $http->post(getApiUrl("/products"), [ 'json' => $data ]);
        }

        $data = $response->getBody();
        $data = json_decode($data->getContents());
        $flashKey = $isEdit ? 'productUpdate' : 'productCreate';
        $action = $isEdit ? 'atualizado' : 'adicionado';
        session()->flash($flashKey, "Produto $action com sucesso");

        return redirect()->route('products');
    }

    public function delete(string $id)
    {
        $http = new RequestsController();

        $response = $http->delete(getApiUrl("/products/$id"));
        $data = $response->getBody();
        $data = json_decode($data->getContents());
        session()->flash('productDelete', "Produto excluído com sucesso");

        return redirect()->route('products')->with('data', $data);
    }
}
