import Fastify from "fastify";
import sharp from "sharp";

const app = new Fastify({ logger: true });

app.get("/img/*", async (request, reply) => {
  const imageUrl = request.params["*"];
  const imgRes = await fetch(imageUrl);
  const imgBuffer = await imgRes.arrayBuffer();
  console.log(request.query);
  console.log({ imageUrl });

  const optimizedImage = await sharp(imgBuffer).resize(600).webp().toBuffer();

  reply.type('image/webp')
  reply.send(optimizedImage)
});

try {
  await app.listen({ port: 3000 });
} catch (error) {
  app.log.error(error);
  process.exit(1);
}
