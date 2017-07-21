export default function gettext(str) {
  return {
    Home: 'Acasã',
    Sales: 'Vânzãri',
    Expenses: 'Cheltuieli',
    Bank: 'Bancã',
    Clients: 'Clienți',
    Invoices: 'Facturi',
    'Add invoice': 'Adaugã',
    Date: 'Data',
    Series: 'Serie',
    Client: 'Client',
    Amount: 'Suma',
    Authentication: 'Autentificare',
    Login: 'Logare',
    Password: 'Parolã',
    'Remember me': 'Ține-mã minte',
    'Forgot my password': 'Am uitat parola',
    'Due by': 'Scadențã',
    Paid: 'Achitat',
  }[str] || str;
}
