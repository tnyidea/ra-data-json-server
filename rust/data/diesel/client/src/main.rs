use std::env;

use diesel::pg::PgConnection;
use diesel::prelude::*;
use dotenv::dotenv;

use diesel_model::address::*;


fn main() {
    use diesel_model::schema::addresses::dsl::*;

    // let connection = &mut establish_connection();

    dotenv().ok();

    let database_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set");
    let connection = &mut PgConnection::establish(&database_url)
        .unwrap_or_else(|_| panic!("Error connecting to {}", database_url));

    let results = addresses
        .limit(5)
        .load::<Address>(connection)
        .expect("Error loading addresses");

    println!("Displaying {} addresses", results.len());
    for item in results {
        println!("{}", item.id);
        println!("-----------\n");
        println!("{}", item.first_name);
        println!("{}", item.last_name);
    }
}