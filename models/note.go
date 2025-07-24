package models

//This declares that the file belongs to the models package in your Go project.
import "gorm.io/gorm"

/*✅ Why we do this:
In Go, every file must belong to a package.

We create a separate models package to keep database-related code (like struct definitions) organized and separate from other parts of the app.
*/

type Note struct {
	gorm.Model        // 🔧 gorm.Model: Embeds ID, CreatedAt, UpdatedAt, DeletedAt fields automatically
	Title      string `json:"title"`
	Content    string `json:"content"`
}

/*🔧 gorm.Model
✅ Meaning:
This is a predefined struct in GORM that you're embedding into your Note struct.

✅ Why we do this:
gorm.Model adds 4 built-in fields automatically:

go
Copy
Edit
ID        uint           // Primary Key (auto-increment)
CreatedAt time.Time      // Auto-set when record is created
UpdatedAt time.Time      // Auto-updated on changes
DeletedAt gorm.DeletedAt // Used for soft deletes
✅ Helps with:
You don’t need to manually define ID, CreatedAt, etc.

GORM takes care of tracking creation/update times and record identity.

*/
