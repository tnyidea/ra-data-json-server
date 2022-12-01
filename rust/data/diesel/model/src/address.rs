use diesel::prelude::*;

#[derive(Queryable)]
pub struct Address {
    pub id: String,
    // pub created_at: PgTimestamp,
    // pub updated_at: PgTimestamp,
    // pub deleted_at: PgTimestamp,
    pub first_name: String,
    pub last_name: String,
    pub company_name: String,
    pub address: String,
    pub city: String,
    pub county: String,
    pub state: String,
    pub zip: String,
    pub phone_1: String,
    pub phone_2: String,
    pub email: String,
    pub web: String,
}