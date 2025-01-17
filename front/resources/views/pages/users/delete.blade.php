@extends('.template.default')

@section('body')
    @include('components.alert', ['flashKey' => 'userDelete'])

    <h1>Deseja mesmo excluir o usuário: {{ $data->Name }}?</h1>

    <form action="{{ route('user.delete', ['id' => $data->ID])}}" method="post">
        @csrf()
        @method('DELETE')

        <a href="{{ route('users') }}" type="button" class="btn btn-secondary">Cancelar</a>
        <button type="submit" class="btn btn-primary">Excluir</button>
    </form>
@endsection
