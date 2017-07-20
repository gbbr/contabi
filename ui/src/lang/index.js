export default function gettext(str) {
  return {
    'Home': 'Acasã',
    'Sales': 'Vânzãri',
    'Expenses': 'Cheltuieli',
    'Bank': 'Banca',
    'Clients': 'Clienți',
    'Invoices': 'Facturi',
  }[str] || str;
}
