package pages

import (
	"github.com/a-h/templ"
	"ostkost/go-ps-hw-fiber/views"
	"ostkost/go-ps-hw-fiber/views/components"
)

func IndexComponent() templ.Component {
	links := []components.NavbarItemProps{
		{Href: "/", Text: "Eда", Img: "food.jpeg"},
		{Href: "/", Text: "Животные", Img: "animals.jpeg"},
		{Href: "/", Text: "Машины", Img: "cars.jpeg"},
		{Href: "/", Text: "Спорт", Img: "sport.jpeg"},
		{Href: "/", Text: "Музыка", Img: "music.jpeg"},
		{Href: "/", Text: "Технологии", Img: "technology.jpeg"},
		{Href: "/", Text: "Прочее", Img: "other.jpeg"},
	}

	postsTitle := "Популярное"
	posts := []components.PostCardProps{
		{Title: "Открытие сезона байдарок", Description: "Сегодня был открыт сезон путешествия на байдарках, где вы можете поучаствовать в ...", Img: "nature.jpg", Username: "Михаил Аршинов", AvatarImg: "Mike.jpg", Date: "Август 18 , 2025"},
		{Title: "Выбери правильный ноутбук для задач", Description: "От верного выбора ноутбука зависит не только удобство, но и эффективность работы...", Img: "mac.jpg", Username: "Вася Программист", AvatarImg: "Vasya.jpg", Date: "Июль 25 , 2025"},
		{Title: "Создание автомобилей с автопилотом", Description: "Электические автомобили без водителя скоро станут реальностью, где нам не придётся ...", Img: "car.jpg", Username: "Мария", AvatarImg: "Mary.jpg", Date: "Июль 14 , 2025"},
		{Title: "Как быстро приготовить вкусный обед", Description: "Сегодня поговорим о том, как можно быстро и эффективно приготовить обед для ...", Img: "food.jpg", Username: "Ли Сюн", AvatarImg: "Li.jpg", Date: "Май 10 , 2025"},
	}

	news := []components.NewsCardProps{
		{Title: "Как безопасно водить", Description: "Длинный текст про то, как можно безопасно водить автомобиль.", Img: "car.jpg"},
		{Title: "Создавай музыку!", Description: "Сегодня мы рассмотрим технику быстрого создания музыки за счёт использования...", Img: "music.jpg"},
	}

	newsSlides := []components.NewsSlideProps{
		{Title: "Несколько мониторов - Зло!", Description: "Большинство людей используют несколько мониторов. Сегодня мы разберём почему это может быть очень не эффективно и как с этим боро...", Img: "monitor.jpg"},
	}

	props := views.IndexProps{
		NavItems:   links,
		PostItems:  posts,
		PostsTitle: postsTitle,
		News:       news,
		NewsSlides: newsSlides,
	}

	return views.Index(props)
}
