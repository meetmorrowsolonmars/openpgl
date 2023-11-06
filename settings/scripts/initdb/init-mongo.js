db.createUser({
    user: process.env.MONGODB_USER,
    pwd: process.env.MONGODB_PASSWORD,
    roles: ["readWrite"],
});

db.createCollection("pallets_settings", {
    clusteredIndex: {
        key: {_id: 1},
        unique: true,
    },
    validator: {
        $jsonSchema: {
            bsonType: "object",
            title: "Pallets Settings Object Validation",
            required: ["userId", "pallets"],
            properties: {
                userId: {
                    bsonType: "int",
                    minimum: 1,
                    description: "'userId' must be greater than zero",
                },
                pallets: {
                    bsonType: "array",
                    maxItems: 5,
                }
            },
        },
    },
});

db.pallets_settings.createIndex({
    userId: 1,
}, {
    unique: true,
});

db.createCollection("default_pallets_settings", {
    clusteredIndex: {
        key: {_id: 1},
        unique: true,
    },
});
