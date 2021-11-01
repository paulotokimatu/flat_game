#version 410

in vec4 vertex; // <vec2 position, vec2 texCoords>
out vec2 TexCoords; //pass to frag

uniform vec2 resolution; //window res

void main()
{
  vec2 zeroToOne = vertex.xy / resolution; // convert the rectangle from pixels to 0.0 to 1.0
  vec2 zeroToTwo = zeroToOne * 2.0; // convert from 0->1 to 0->2
  vec2 clipSpace = zeroToTwo - 1.0; // convert from 0->2 to -1->+1 (clipspace)
  gl_Position = vec4(clipSpace * vec2(1, -1), 0, 1);
  TexCoords = vertex.zw;
}
