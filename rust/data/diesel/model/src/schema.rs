// @generated automatically by Diesel CLI.

diesel::table! {
    addresses (id) {
        id -> Varchar,
        first_name -> Varchar,
        last_name -> Varchar,
        company_name -> Varchar,
        address -> Varchar,
        city -> Varchar,
        county -> Varchar,
        state -> Varchar,
        zip -> Varchar,
        phone1 -> Varchar,
        phone2 -> Varchar,
        email -> Varchar,
        web -> Varchar,
    }
}
