import type { User } from "./user_types";

export class Post {
    public id: string;
    public createdAt: Date;
    public updatedAt: Date;
    public caption: string;
    public images: string[];
    public poster: User;

    constructor(id: string, createdAt: Date, updatedAt: Date, caption: string, images: string[], poster: User) {
        this.id = id;
        this.createdAt = createdAt;
        this.updatedAt = updatedAt;
        this.caption = caption;
        this.images = images;
        this.poster = poster;
    }
}