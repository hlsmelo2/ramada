@extends('.template.default')

@section('body')
    @php
        $title = $mode === 'create' ? 'Novo produto' : 'Atualizar produto';
        $buttonTitle = $mode === 'create' ? 'Adicionar' : 'Atualizar';
        $method = $mode === 'create' ? 'POST' : 'PUT';
        $route = $mode === 'create' ? 'product.create' : 'product.edit';

        $fields = [
            ['Nome', $data->Name, 'name', 'text'],
            ['Descrição', $data->Description, 'description', 'text'],
            ['Preço', currency($data->Price), 'price', 'text'],
            ['Categoria', $data->Category, 'category', 'text'],
        ];
    @endphp

    <h1>{{ $title }}</h1>

    <form action="{{ route('product.edit', ['id' => $data->ID])}}" method="post">
        @csrf()
        @method($method)

        @foreach ($fields as $field)
            <div class="input-group mb-3">
                <span class="input-group-text" id="{{ $field[1] }}">{{ $field[0] }}</span>
                <input type="{{ $field[3] }}" name="{{ $field[2] }}" class="form-control" value="{{ $field[1] }}" placeholder="{{ $field[0] }}" aria-label="{{ $field[1] }}" aria-describedby="basic-addon1">
            </div>
        @endforeach

        <a href="{{ route('products') }}" type="button" class="btn btn-secondary">Cancelar</a>
        <button type="submit" class="btn btn-primary">{{ $buttonTitle }}</button>
    </form>
@endsection
