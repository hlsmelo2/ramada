@extends('.template.default')

@section('body')
    @include('components.alert', ['flashKey' => 'productDelete'])

    <h1>Deseja mesmo excluir o produto: {{ $data->Name }}?</h1>

    <form action="{{ route('product.delete', ['id' => $data->ID])}}" method="post">
        @csrf()
        @method('DELETE')

        <a href="{{ route('products') }}" type="button" class="btn btn-secondary">Cancelar</a>
        <button type="submit" class="btn btn-primary">Excluir</button>
    </form>
@endsection
