BEGIN;

CREATE TABLE IF NOT EXISTS borrowing_records (
    borrow_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    book_id UUID NOT NULL,
    borrowed_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    returned_at TIMESTAMPTZ, 
    
   CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(user_id),
    CONSTRAINT fk_book FOREIGN KEY (book_id) REFERENCES books(book_id)
);

COMMIT;
