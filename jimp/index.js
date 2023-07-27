import Fastify from "fastify";
import Jimp from "jimp";

const app = new Fastify({ logger: true });

app.get("/img/*", async (request, reply) => {
  const imageUrl = request.params["*"];

  const originalImage = await Jimp.read(imageUrl);
  originalImage.resize(600, Jimp.AUTO).getBuffer(Jimp.MIME_JPEG, (err, buffer) => {
    if (err) {
      throw new Error(err);
    }
    reply.type("image/jpeg");
    reply.send(buffer);
  });
});

try {
  await app.listen({ port: 3000 });
} catch (error) {
  app.log.error(error);
  process.exit(1);
}
