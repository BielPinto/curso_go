mutation createCategory {
  createCategory(input: {name: "Biologia", description: "Cursos de Medicina"}) {
    id
    name
    description
  }
}

mutation createCourse {
  createCourse(
    input: {name: "Full cycle", description: "the best", categoryId: "cd0360b6-a529-4045-bccc-6681ca226750"}
  ) {
    id
    name
  }
}

query queryCategories {
  categories {
    id
    name
    description
  }
}

query queryCategoriesWithCourses {
  categories {
    id
    name
    description
    courses {
      id
      name
    }
  }
}

query queryCourses {
  courses {
    id
    name
  }
}

query queryCoursesWithCategory {
  courses {
    id
    name
    description
    category {
      id
      name
      description
    }
  }
}