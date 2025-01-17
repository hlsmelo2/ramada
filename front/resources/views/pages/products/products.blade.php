@extends('.template.default')

@section('body')
    @include('components.alert', ['flashKey' => 'productCreate'])
    @include('components.alert', ['flashKey' => 'productUpdate'])
    @include('components.alert', ['flashKey' => 'productDelete'])

    <h1>Produtos</h1>

    <form style="margin-bottom: 30px;" action="">
        <div class="input-group mb-3">
            <label for="nome" class="input-group-text">Nome</label>
            <input id="nome" name="nome" class="form-control" type="text" value="{{ $filters['nome'] }}">
        </div>

        @include('components.categories')

        <div class="input-group mb-3">
            <label for="preco_min" class="input-group-text">Preço Minimo</label>
            <input id="preco_min" name="preco_min" class="form-control" type="text" value="{{ $filters['preco_min'] }}">
        </div>

        <div class="input-group mb-3">
            <label for="preco_max" class="input-group-text">Preço Máximo</label>
            <input id="preco_max" name="preco_max" class="form-control" type="text" value="{{ $filters['preco_max'] }}">
        </div>

        <input type="submit" class="btn btn-secondary" value="Filtrar">
    </form>

    <div class="table-responsive">
        <table class="table">
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Nome</th>
                    <th>Descrição</th>
                    <th>Preço</th>
                    <th>Categoria</th>
                    <th>Ações</th>
                </tr>
            </thead>
            <tbody>
                @foreach ($data as $item)
                <tr>
                    <td>{{ $item->ID}}</td>
                    <td>{{ $item->Name}}</td>
                    <td>{{ $item->Description}}</td>
                    <td>R$ {{ currency($item->Price) }}</td>
                    <td>{{ $item->Category}}</td>
                    <td>
                        <a href="{{ route('page.product.edit', ['id' => $item->ID]) }}" class="edit">Editar</a>
                        <a href="{{ route('page.product.delete', ['id' => $item->ID]) }}" class="delete">Excluir</a>
                    </td>
                </tr>
                @endforeach
            </tbody>
        </table>
    </div>
@endsection