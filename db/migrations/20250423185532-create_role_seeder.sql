
-- +migrate Up

INSERT INTO "role"(
    "id",
    "name",
    "created_at",
    "updated_at"
) VALUES
(
    '8d5a466e-9342-4330-8226-d67de9d82563',
    'borrower',
    NOW(),
    NOW()
),
(
    '9b5c5d13-919d-45e0-b07e-3aec8d8d5fb8',
    'investor',
    NOW(),
    NOW()
),
(
    '62d78853-8f3a-46f1-aa3d-5b4c0fc294fb',
    'employee',
    NOW(),
    NOW()
);

-- +migrate Down

DELETE FROM "role" WHERE "id" IN ('8d5a466e-9342-4330-8226-d67de9d82563', '9b5c5d13-919d-45e0-b07e-3aec8d8d5fb8', '62d78853-8f3a-46f1-aa3d-5b4c0fc294fb');
 