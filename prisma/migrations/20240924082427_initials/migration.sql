-- CreateTable
CREATE TABLE "Post" (
    "id" VARCHAR(100) NOT NULL DEFAULT gen_random_uuid(),
    "created_at" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3),
    "title" VARCHAR(100) NOT NULL,
    "published" BOOLEAN DEFAULT false,
    "description" VARCHAR(100) NOT NULL,

    CONSTRAINT "Post_pkey" PRIMARY KEY ("id")
);
