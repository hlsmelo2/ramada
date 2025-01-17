<?php

function isAuthed() {
    return ($_COOKIE['_token'] ?? null) !== null;
}

function getApiUrl(string $endpoint): string {
    return env('API_URL') . $endpoint;
}

function currency(float $value): string {
    return number_format($value, 2, ',', '.');
}

function currencyFloat(string $currency) {
    $currency = preg_replace("#,+#", "&&", $currency);
    $currency = preg_replace("#\.+#", ",", $currency);
    $currency = preg_replace("#&&+#", ".", $currency);

    return $currency;
}
