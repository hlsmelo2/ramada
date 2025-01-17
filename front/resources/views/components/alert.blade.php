@php
    $typeClasses = [
        'success' => 'success alert-success',
        'danger' => 'danger alert-danger',
    ][$type ?? 'success'];
@endphp

@if (session()->has($flashKey))
    <div class="alert {{ $typeClasses }}">{{ session()->get($flashKey) }}</div>
@endif
