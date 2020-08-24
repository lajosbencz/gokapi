
export default () => ({
  list: [
    {
      id: 0,
      name: 'Block - Simple',
      fields: [
        {
          name: 'title',
          type: 'string',
          required: true,
        },
        {
          name: 'content',
          type: 'html',
          required: true,
        },
      ],
    },
  ],
})
