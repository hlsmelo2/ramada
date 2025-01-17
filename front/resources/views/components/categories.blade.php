@php
    $categories = [
        [ 'label' => 'Categoria 1' , 'value' => 'Category 1' ],
        [ 'label' => 'Categoria 2' , 'value' => 'Category 2' ],
        [ 'label' => 'Categoria 3' , 'value' => 'Category 3' ],
        [ 'label' => 'Categoria 4' , 'value' => 'Category 4' ],
    ];
@endphp

<div class="input-group mb-3">
    <label for="categoria" class="input-group-text">Categoria</label>
    <select name="categoria" id="categoria" class="form-control" value="{{ $filters['categoria'] }}">
        @foreach ($categories as $category)
            @if ($filters['categoria'] === $category['value'])
                <option selected value="{{ $category['value'] }}">{{ $category['label'] }}</option>
            @else
                <option value="{{ $category['value'] }}">{{ $category['label'] }}</option>
            @endif
        @endforeach
    </select>
</div>
