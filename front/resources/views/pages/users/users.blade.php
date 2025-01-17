@extends('.template.default')

@section('body')
    @include('components.alert', ['flashKey' => 'userCreate'])
    @include('components.alert', ['flashKey' => 'userUpdate'])
    @include('components.alert', ['flashKey' => 'userDelete'])

    <h1>Usuários</h1>

    <div class="table-responsive">
        <table class="table">
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Nome</th>
                    <th>Nome de usuário</th>
                    <th>Email</th>
                    <th>Ações</th>
                </tr>
            </thead>
            <tbody>
                @foreach ($data as $item)
                <tr>
                    <td>{{ $item->ID}}</td>
                    <td>{{ $item->Name}}</td>
                    <td>{{ $item->Username}}</td>
                    <td>{{ $item->Email}}</td>
                    <td>
                        <a href="{{ route('page.user.edit', ['id' => $item->ID]) }}" class="edit">Editar</a>
                        <a href="{{ route('page.user.delete', ['id' => $item->ID]) }}" class="delete">Excluir</a>
                    </td>
                </tr>
                @endforeach
            </tbody>
        </table>
    </div>
@endsection