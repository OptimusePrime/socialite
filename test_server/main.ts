import cors from "@fastify/cors";
import Fastify from "fastify";

const fastify = Fastify({
    logger: true,
})
fastify.register(cors)

fastify.get("/test", function (req, res) {
    res.send({
        hello: "world"
    })
})

fastify.listen({ port: 5000 }, function (err, address) {
    if (err) {
        fastify.log.error(err)
        process.exit(1)
    }
})