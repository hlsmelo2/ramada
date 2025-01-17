@extends('.template.default')

@section('body')
    @php
        $title = $mode === 'create' ? 'Novo usuário' : 'Atualizar usuário';
        $buttonTitle = $mode === 'create' ? 'Adicionar' : 'Atualizar';
        $method = $mode === 'create' ? 'POST' : 'PUT';
        $route = $mode === 'create' ? route('user.create') : route('user.edit', ['id' => $data->ID]);

        $fields = [
            ['Nome', $data->Name, 'name', 'text'],
            ['Nome de usuário', $data->Username, 'username', 'text'],
            ['Email', $data->Email, 'email', 'text'],
            ['Senha', '', 'password', 'password'],
            ['Repetir senha', '', 'repassowrd', 'password'],
        ];
    @endphp

    <h1>{{ $title }}</h1>

    <form action="{{ $route }}" method="post">
        @csrf()
        @method($method)

        @foreach ($fields as $field)
            <div class="input-group mb-3">
                <span class="input-group-text" id="{{ $field[1] }}">{{ $field[0] }}</span>
                <input type="{{ $field[3] }}" name="{{ $field[2] }}" class="form-control" value="{{ $field[1] }}" placeholder="{{ $field[0] }}" aria-label="{{ $field[1] }}" aria-describedby="basic-addon1">
            </div>
        @endforeach

        <a href="{{ route('users') }}" type="button" class="btn btn-secondary">Cancelar</a>
        <button type="submit" class="btn btn-primary">{{ $buttonTitle }}</button>
    </form>
@endsection
